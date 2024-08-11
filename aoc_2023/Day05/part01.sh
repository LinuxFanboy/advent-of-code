#!/bin/bash

map_number() {
    local number=$1
    shift
    while [[ $# -gt 0 ]]; do
        local dest_start=$1
        local src_start=$2
        local length=$3
        shift 3
        if [[ $number -ge $src_start && $number -lt $(($src_start + $length)) ]]; then
            local offset=$(($number - $src_start))
            echo $(($dest_start + offset))
            return
        fi
    done
    echo $number
}

process_seed() {
    local seed=$1
    shift
    local number=$seed
    for map in "$@"; do
        number=$(map_number $number $map)
    done
    echo $number
}

file_path="input" # Adjust the file path as necessary
readarray -t input_data < $file_path

IFS=' ' read -ra seed_numbers <<< "${input_data[0]#*: }"

mappings=()
current_map=()
for line in "${input_data[@]:2}"; do
    if [[ $line == *"map:"* ]]; then
        if [[ ${#current_map[@]} -gt 0 ]]; then
            mappings+=("$(IFS=$' '; echo "${current_map[*]}")")
            current_map=()
        fi
    else
        current_map+=($line)
    fi
done
if [[ ${#current_map[@]} -gt 0 ]]; then
    mappings+=("$(IFS=$' '; echo "${current_map[*]}")")
fi

lowest_location_number=0
for seed in "${seed_numbers[@]}"; do
    location_number=$(process_seed $seed "${mappings[@]}")
    if [[ $lowest_location_number -eq 0 || $location_number -lt $lowest_location_number ]]; then
        lowest_location_number=$location_number
    fi
done

echo "Lowest location number: $lowest_location_number"

