#!/usr/bin/env ruby

total_points=0

ARGF.each do |line|
  winning_numbers, my_numbers = line.split(': ')[1].split(' | ').map { |s| s.split.map(&:to_i) }
  matches = winning_numbers & my_numbers
  total_points += matches.empty? ? 0 : 2 ** (matches.size - 1)
end

puts "Total points: #{total_points}"