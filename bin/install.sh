#!/bin/bash

# Build the Go program
go build -o vodm cmd/vodm/main.go

# Create bin directory if it doesn't exist
mkdir -p ~/bin

# Move the binary to bin directory
mv vodm ~/bin/

# Add ~/bin to PATH if not already there
if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
    echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
    echo "Added ~/bin to PATH. Please restart your terminal or run: source ~/.zshrc"
fi

echo "Installation complete! You can now use 'vodm' from anywhere."
