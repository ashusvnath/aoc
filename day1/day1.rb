#!/usr/bin/env ruby
require 'logger'
logger = Logger.new(STDERR)
logger.level = Logger::ERROR
require 'optparse'
input = nil

OptionParser.new do |opts|
  opts.banner = "Usage: #{$0} [options]"

  opts.on("-v", "--[no-]verbose", "Run verbosely") do |v|
    logger.level = v ? Logger::DEBUG : Logger::ERROR
  end

  opts.on("-f [FILE]", "--file [FILE_PATH]", "Path to input file ") do |file_path|
    input = file_path
  end
end.parse!

if input == nil
    print "Name of file: "
    input = STDIN.gets.chomp
end

logger.debug "Input file:" + input

all_elf_data = {}
part2_max = []
num_elves = 1
current_elf_calories = []

File.read(input).lines.each do |line|
    calories = line.chomp.to_i
    current_elf_calories << calories
    if calories == 0
        elf_data = { calories: current_elf_calories, total: current_elf_calories.sum }
        logger.debug "#{num_elves}, #{elf_data[:total]}"
        all_elf_data[num_elves] = elf_data
        num_elves += 1
        current_elf_calories = []
        part2_max << elf_data[:total]
        part2_max = part2_max.sort{|x,y| y <=> x}.take(3)
        logger.debug "Part1 max so far: #{part2_max[0]}"
        logger.debug "Part2 max so far: #{part2_max}"
    end
end
puts "Part 1 max: #{part2_max[0]}"
puts "Part 2 max: #{part2_max.sum}"