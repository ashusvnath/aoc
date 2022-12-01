require 'logger'
logger = Logger.new(STDERR)
logger.level = Logger::ERROR
require 'optparse'
input = nil

OptionParser.new do |opts|
  opts.banner = "Usage: #{$0} [options]"

  opts.on("-v", "--[no-]verbose", "Run verbosely") do
    logger.level = Logger::DEBUG
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


data = File.read(input).lines.map(&:chomp)
all_elf_data = {}
part1_max = 0
part2_max = []
num_elves = 1
current_calories = []
result = data.reduce { |elves, l|
    if l != ""
        current_calories << l.to_i
    else
        elf_data = { calories: current_calories, total: current_calories.sum }
        logger.debug "#{num_elves}, #{elf_data[:total]}"
        all_elf_data[num_elves] = elf_data
        num_elves += 1
        current_calories = []
        part1_max = elf_data[:total] > part1_max ? elf_data[:total] : part1_max
        part2_max << elf_data[:total]
        part2_max = part2_max.sort.reverse.take(3)
    end
    all_elf_data
}
puts "Part 1 max: #{part1_max}"
puts "Part 2 max: #{part2_max.sum}"