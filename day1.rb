require 'logger'
logger = Logger.new(STDERR)

input = nil
if ARGV.length > 0
    input = ARGV[0]
else
    logger.debug "Name of file: "
    input = STDIN.gets.chomp
end
logger.debug "Input:" + input


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