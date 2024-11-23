#!/bin/bash

# Ensure the script stops on errors
set -e

# Default commit message
DEFAULT_MESSAGE="Update made on $(date)"

# Check for commit message argument
if [ -z "$1" ]; then
  echo "No commit message provided. Using default: '$Go'"
  COMMIT_MESSAGE="$DEFAULT_MESSAGE"
else
  COMMIT_MESSAGE="$1"
fi

# Add all changes
echo "Adding changes..."
git add .

# Commit changes
echo "Committing changes with message: $COMMIT_MESSAGE"
git commit -m "$COMMIT_MESSAGE"

# Push changes to the remote repository
echo "Pushing changes to the remote repository..."
git push

echo "All done!"
