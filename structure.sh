#!/bin/bash

if [ ! -d "./dsa" ]; then
    echo "Error: './dsa' directory does not exist in the current location."
    exit 1
fi

echo "Scanning categories inside ./dsa..."
echo "-----------------------------------"

dirs=()
while IFS= read -r dir; do
    dirs+=("$dir")
done < <(find ./dsa -mindepth 1 -maxdepth 1 -type d | sort)

if [ ${#dirs[@]} -eq 0 ]; then
    echo "No directories found inside ./dsa. Create one first!"
    exit 1
fi

for i in "${!dirs[@]}"; do
    folder_name="${dirs[$i]#./dsa/}"
    echo "$((i + 1))) $folder_name"
done

echo "-----------------------------------"
read -p "Select a category number: " choice

if ! [[ "$choice" =~ ^[0-9]+$ ]] || [ "$choice" -lt 1 ] || [ "$choice" -gt "${#dirs[@]}" ]; then
    echo "Error: Invalid selection. Please run the script again."
    exit 1
fi

selected_category="${dirs[$((choice - 1))]}"

read -p "Enter the new problem folder name (e.g., concatenation_of_array): " problem_name

if [ -z "$problem_name" ]; then
    echo "Error: Problem name cannot be empty."
    exit 1
fi

target_dir="$selected_category/$problem_name"

if [ -d "$target_dir" ]; then
    echo "Error: '$target_dir' already exists!"
    exit 1
fi

echo "Creating structure at $target_dir..."
mkdir -p "$target_dir"

cd "$target_dir" || exit 1

touch better_solution.md
touch discussion.md
touch main.go
touch main_test.go
touch main_bench_test.go
touch question.md

echo "Successfully created all files inside $target_dir!"