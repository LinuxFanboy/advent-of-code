#!/usr/bin/env ruby

schematic = ARGF.readlines.map { |line| ".#{line.chomp}." }
schematic.unshift('.' * schematic[0].length).push('.' * schematic[0].length)

part_sum = 0
gears = Hash.new { |hash, key| hash[key] = [] }

(1..(schematic.size - 2)).each do |line_index|
  schematic[line_index].scan(/\d+/) do |number|
    x_start = $~.offset(0)[0] - 1

    found = false
    ((line_index - 1)..(line_index + 1)).each do |y|
      next if x_start.negative? || x_start + number.length + 2 > schematic[y].length
      substring_to_check = schematic[y][x_start, number.length + 2]
      found = true if substring_to_check&.match?(/[^.\d]/)
      substring_to_check.scan(/\*/) { gears[[y, x_start + $~.offset(0)[0]]] << number.to_i }
    end

    part_sum += number.to_i if found
  end
end

part_2_sum = gears.values.select { |values| values.size == 2 }.sum { |values| values.reduce(:*) }

puts "Part 1: #{part_sum}"
puts "Part 2: #{part_2_sum}"
