package lassie

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfsnode"
	"github.com/ipfs/go-unixfsnode/data"
	"github.com/ipfs/go-unixfsnode/file"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

var ErrNotDir = fmt.Errorf("not a directory")

func pathSegments(path string) ([]string, error) {
	segments := strings.Split(path, "/")
	filtered := make([]string, 0, len(segments))
	for i := 0; i < len(segments); i++ {
		if segments[i] == "" {
			// Allow one leading and one trailing '/' at most
			if i == 0 || i == len(segments)-1 {
				continue
			}
			return nil, fmt.Errorf("invalid empty path segment at position %d", i)
		}
		if segments[i] == "." || segments[i] == ".." {
			return nil, fmt.Errorf("'%s' is unsupported in paths", segments[i])
		}
		filtered = append(filtered, segments[i])
	}
	return filtered, nil
}

func extractRoot(ls *ipld.LinkSystem, root cid.Cid, outputDir string, path []string) (int, error) {
	if root.Prefix().Codec == cid.Raw {
		// if c.IsSet("verbose") {
		// 	fmt.Fprintf(c.App.ErrWriter, "skipping raw root %s\n", root)
		// }
		return 0, nil
	}

	pbn, err := ls.Load(ipld.LinkContext{}, cidlink.Link{Cid: root}, dagpb.Type.PBNode)
	if err != nil {
		return 0, err
	}
	pbnode := pbn.(dagpb.PBNode)

	ufn, err := unixfsnode.Reify(ipld.LinkContext{}, pbnode, ls)
	if err != nil {
		return 0, err
	}

	var outputResolvedDir string
	if outputDir != "-" {
		//log.Println("output dir", outputDir)
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			if err := os.Mkdir(outputDir, 0755); err != nil {
				return 0, err
			}
		}
		outputResolvedDir, err = filepath.EvalSymlinks(outputDir)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		//log.Println("outputResolvedDir", outputResolvedDir)
		if _, err := os.Stat(outputResolvedDir); os.IsNotExist(err) {
			if err := os.Mkdir(outputResolvedDir, 0755); err != nil {
				return 0, err
			}
		}
	}

	//log.Println(outputResolvedDir, path)

	count, err := extractDir(ls, ufn, outputResolvedDir, "/", path)
	if err != nil {
		if !errors.Is(err, ErrNotDir) {
			return 0, fmt.Errorf("%s: %w", root, err)
		}

		// if it's not a directory, it's a file.
		ufsData, err := pbnode.LookupByString("Data")
		if err != nil {
			return 0, err
		}
		ufsBytes, err := ufsData.AsBytes()
		if err != nil {
			return 0, err
		}
		ufsNode, err := data.DecodeUnixFSData(ufsBytes)
		if err != nil {
			return 0, err
		}
		var outputName string
		if outputDir != "-" {
			outputName = filepath.Join(outputResolvedDir, "unknown")
		}

		//log.Println("outputResolvedDir", outputResolvedDir)
		//log.Println("outputName", outputName)
		if ufsNode.DataType.Int() == data.Data_File || ufsNode.DataType.Int() == data.Data_Raw {
			if err := extractFile(ls, pbnode, outputName); err != nil {
				return 0, err
			}
		}
		return 1, nil
	}

	return count, nil
}
func resolvePath(root, pth string) (string, error) {
	rp, err := filepath.Rel("/", pth)
	if err != nil {
		return "", fmt.Errorf("couldn't check relative-ness of %s: %w", pth, err)
	}
	joined := path.Join(root, rp)

	basename := path.Dir(joined)
	final, err := filepath.EvalSymlinks(basename)
	if err != nil {
		return "", fmt.Errorf("couldn't eval symlinks in %s: %w", basename, err)
	}
	if final != path.Clean(basename) {
		return "", fmt.Errorf("path attempts to redirect through symlinks")
	}
	return joined, nil
}

func extractDir(ls *ipld.LinkSystem, n ipld.Node, outputRoot, outputPath string, matchPath []string) (int, error) {
	if outputRoot != "" {
		dirPath, err := resolvePath(outputRoot, outputPath)
		if err != nil {
			return 0, err
		}
		// make the directory.
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return 0, err
		}
	}

	if n.Kind() != ipld.Kind_Map {
		return 0, ErrNotDir
	}

	subPath := matchPath
	if len(matchPath) > 0 {
		subPath = matchPath[1:]
	}

	extractElement := func(name string, n ipld.Node) (int, error) {
		var nextRes string
		if outputRoot != "" {
			var err error
			nextRes, err = resolvePath(outputRoot, path.Join(outputPath, name))
			if err != nil {
				return 0, err
			}
			// if c.IsSet("verbose") {
			// 	fmt.Fprintf(c.App.ErrWriter, "%s\n", nextRes)
			// }
		}

		if n.Kind() != ipld.Kind_Link {
			return 0, fmt.Errorf("unexpected map value for %s at %s", name, outputPath)
		}
		// a directory may be represented as a map of name:<link> if unixADL is applied
		vl, err := n.AsLink()
		if err != nil {
			return 0, err
		}
		dest, err := ls.Load(ipld.LinkContext{}, vl, basicnode.Prototype.Any)
		if err != nil {
			if nf, ok := err.(interface{ NotFound() bool }); ok && nf.NotFound() {
				fmt.Printf("data for entry not found: %s (skipping...)\n", path.Join(outputPath, name))
				return 0, nil
			}
			return 0, err
		}
		// degenerate files are handled here.
		if dest.Kind() == ipld.Kind_Bytes {
			if err := extractFile(ls, dest, nextRes); err != nil {
				return 0, err
			}
			return 1, nil
		}

		// dir / pbnode
		pbb := dagpb.Type.PBNode.NewBuilder()
		if err := pbb.AssignNode(dest); err != nil {
			return 0, err
		}
		pbnode := pbb.Build().(dagpb.PBNode)

		// interpret dagpb 'data' as unixfs data and look at type.
		ufsData, err := pbnode.LookupByString("Data")
		if err != nil {
			return 0, err
		}
		ufsBytes, err := ufsData.AsBytes()
		if err != nil {
			return 0, err
		}
		ufsNode, err := data.DecodeUnixFSData(ufsBytes)
		if err != nil {
			return 0, err
		}

		switch ufsNode.DataType.Int() {
		case data.Data_Directory, data.Data_HAMTShard:
			ufn, err := unixfsnode.Reify(ipld.LinkContext{}, pbnode, ls)
			if err != nil {
				return 0, err
			}
			return extractDir(ls, ufn, outputRoot, path.Join(outputPath, name), subPath)
		case data.Data_File, data.Data_Raw:
			if err := extractFile(ls, pbnode, nextRes); err != nil {
				return 0, err
			}
			return 1, nil
		case data.Data_Symlink:
			if nextRes == "" {
				return 0, fmt.Errorf("cannot extract a symlink to stdout")
			}
			data := ufsNode.Data.Must().Bytes()
			if err := os.Symlink(string(data), nextRes); err != nil {
				return 0, err
			}
			return 1, nil
		default:
			return 0, fmt.Errorf("unknown unixfs type: %d", ufsNode.DataType.Int())
		}
	}

	// specific path segment
	if len(matchPath) > 0 {
		val, err := n.LookupByString(matchPath[0])
		if err != nil {
			return 0, err
		}
		return extractElement(matchPath[0], val)
	}

	if outputPath == "-" && len(matchPath) == 0 {
		return 0, fmt.Errorf("cannot extract a directory to stdout, use a path to extract a specific file")
	}

	// everything
	var count int
	var shardSkip int
	mi := n.MapIterator()
	for !mi.Done() {
		key, val, err := mi.Next()
		if err != nil {
			if nf, ok := err.(interface{ NotFound() bool }); ok && nf.NotFound() {
				shardSkip++
				continue
			}
			return 0, err
		}
		ks, err := key.AsString()
		if err != nil {
			return 0, err
		}
		ecount, err := extractElement(ks, val)
		if err != nil {
			return 0, err
		}
		count += ecount
	}
	if shardSkip > 0 {
		fmt.Printf("data for entry not found for %d unknown sharded entries (skipped...)\n", shardSkip)
	}
	return count, nil
}

func extractFile(ls *ipld.LinkSystem, n ipld.Node, outputName string) error {
	node, err := file.NewUnixFSFile(nil, n, ls)
	if err != nil {
		return err
	}
	nlr, err := node.AsLargeBytes()
	if err != nil {
		return err
	}
	var f *os.File
	if outputName == "" {
		f = os.Stdout
	} else {
		f, err = os.Create(outputName)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	_, err = io.Copy(f, nlr)
	return err
}
