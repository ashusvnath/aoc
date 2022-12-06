#!/usr/bin/env ruby
require 'logger'
logger = Logger.new(STDERR)
logger.level = Logger::ERROR
require 'optparse'
input_filePath = nil

OptionParser.new do |opts|
  opts.banner = "Usage: #{$0} [options]"

  opts.on("-v", "--[no-]verbose", "Run verbosely") do |v|
    logger.level = v ? Logger::DEBUG : Logger::ERROR
  end

  opts.on("-f [FILE]", "--file [FILE_PATH]", "Path to input file ") do |file_path|
    input_filePath = file_path
  end
end.parse!

if input_filePath == nil
    print "Name of file: "
    input_filePath = STDIN.gets.chomp
end

def score(char)
    case char 
    when ('a'..'z')
        score = 1 + char.ord - 'a'.ord
    when ('A'..'Z')
        score = 27 + char.ord - 'A'.ord
    end
end

logger.debug "Input file:" + input_filePath
part1_total = 0
part2_total = 0

File.read(input_filePath).lines.each_slice(3) do |group|
    group.sort_by!(&:length).each do |line|
        left, right = line.chars.each_slice(line.length/2).map(&:join)
        leftAsCharClass = "[#{left}]"
        common = /#{leftAsCharClass}/.match(right)
        line_score = score(common[0])
        part1_total += line_score
        logger.debug "Part1: left: #{left} , right: #{right}, common: #{common}, score: #{line_score}}"    
    end

    smallest = group[0]
    others = group[1..]
    smallestAsCharClass = "[#{smallest}]"
    smallestPattern = /#{smallestAsCharClass}/
    nextPatternString =  group[1].scan(smallestPattern).join("")
    nextPatternStringAsClass = "[#{nextPatternString}]"
    common = /#{nextPatternStringAsClass}/.match(group[2])
    group_score = score(common[0])
    part2_total += group_score
    logger.debug "Part2: #{smallest}, #{nextPatternString}, #{common}, #{group_score}"
end

puts "Part1: #{part1_total}"
puts "Part2: #{part2_total}"
