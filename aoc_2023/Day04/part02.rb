#!/usr/bin/env ruby

def parse_input
  cards = []
  STDIN.each do |line|
    parts = line.strip.split('|')
    card_numbers = parts[0].split(':')[1].strip.split.map(&:to_i)
    winning_numbers = parts[1].strip.split.map(&:to_i)
    cards << { card: card_numbers, win: winning_numbers }
  end
  cards
end

def count_matches(card, winning_numbers)
  (card[:card] & winning_numbers).count
end

def process_cards(cards)
  total_cards = cards.length
  card_copies = Array.new(cards.length, 1)

  cards.each_with_index do |card, index|
    next if index == cards.length - 1

    matches = count_matches(card, card[:win])
    next if matches == 0

    (index + 1..index + matches).each do |i|
      break if i >= cards.length
      card_copies[i] += card_copies[index]
    end
  end

  card_copies.sum
end

cards = parse_input
total_scratchcards = process_cards(cards)
puts "Total scratchcards: #{total_scratchcards}"
