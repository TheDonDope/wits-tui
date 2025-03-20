#!/usr/bin/env bash

target_dir="${1:-.}"

# Validate directory exists
if [[ ! -d "$target_dir" ]]; then
    gum style --foreground 196 --bold "Directory not found: $target_dir"
    exit 1
fi

# Change to target directory
cd "$target_dir" || exit 1

# Find all .tape files
files=(*.tape)
total_files=${#files[@]}

# Check for no files case
if [[ ${total_files} -eq 0 ]]; then
    gum style --foreground 196 --bold "No .tape files found in: $(pwd)"
    exit 1
fi

# Process files with visual feedback
for ((i = 0; i < ${total_files}; i++)); do
    current_file="${files[i]}"
    gum style --bold "Processing $((i + 1))/${total_files}: ${current_file}"
    # Output could be further be reduced by adding >/dev/null 2>&1 to the end
    gum spin --title "Rendering..." -- vhs "${current_file}"
done

# Final message
gum style --foreground 40 --bold "Successfully processed ${total_files} files in: $(pwd)"
