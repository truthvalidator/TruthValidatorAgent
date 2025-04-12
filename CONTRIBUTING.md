# Contributing to TruthValidator

Thank you for your interest in contributing to TruthValidator! We welcome all types of contributions including code, documentation, testing, and community support.

## 🛠️ Development Setup

### Prerequisites
- Go 1.21+
- Docker (recommended)
- Git

### Getting Started
1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/TruthValidator.git
   cd TruthValidator
   ```
3. Set up development environment:
   ```bash
   # Install Go dependencies
   cd VerifyProposalAIAgent
   go mod download
   
   # For smart contract development
   cd contracts/TruthValidatorSentientNet
   make
   ```

## 📝 Code Contribution Guidelines

### Branching Strategy
- `main` - Stable production branch
- `develop` - Integration branch for features
- `feature/*` - Feature development branches
- `bugfix/*` - Bug fix branches

### Commit Messages
Follow [Conventional Commits](https://www.conventionalcommits.org/) format:
```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

Common types:
- feat: New feature
- fix: Bug fix
- docs: Documentation changes
- test: Test related changes
- chore: Maintenance tasks

### Pull Requests
1. Create a feature branch from `develop`
2. Implement your changes with tests
3. Ensure all tests pass:
   ```bash
   go test ./...
   ```
4. Update documentation if needed
5. Submit PR to `develop` branch with:
   - Clear description of changes
   - Related issue number (if applicable)
   - Screenshots (for UI changes)

## 🧪 Testing Requirements
- Unit tests for all new code
- Integration tests for critical paths
- Test coverage should not decrease
- Run tests before submitting PR:
  ```bash
  go test -cover ./...
  ```

## 📜 Code Style
- Go code: Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Solidity: Follow [Solidity Style Guide](https://docs.soliditylang.org/en/latest/style-guide.html)
- Use descriptive variable names
- Add comments for complex logic
- Keep functions focused and small

## 🗣️ Community Guidelines

### Communication
- Be respectful and inclusive
- Use the issue tracker for technical discussions

### Issue Reporting
- Check existing issues before creating new ones
- Use clear, descriptive titles
- Include steps to reproduce for bugs
- Add screenshots when helpful

## 🙏 Thank You!
We appreciate your contributions to building a decentralized truth verification system. Your work helps fight misinformation worldwide.

For any questions, please open an issue in the repository.
