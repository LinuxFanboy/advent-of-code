const std = @import("std");

const ParseError = error{
InvalidFormat,
OutOfMemory,
} || std.fmt.ParseIntError;

fn parseNumber(str: []const u8) ParseError!i64 {
    return std.fmt.parseInt(i64, str, 10) catch |err| switch (err) {
        error.InvalidCharacter => return ParseError.InvalidFormat,
        else => |e| return e,
    };
}

fn is_safe_sequence(numbers: []const i64) bool {
    if (numbers.len < 2) return false;

    var prev = numbers[0];
    var direction: ?bool = null;

    for (numbers[1..]) |num| {
        const diff = num - prev;

        if (diff == 0 or @abs(diff) > 3) {
            return false;
        }

        if (direction) |expected_up| {
            const is_going_up = diff > 0;
            if (expected_up != is_going_up) {
                return false;
            }
        } else {
            direction = diff > 0;
        }

        prev = num;
    }

    return true;
}

fn is_safe_with_dampener(report: []const u8) ParseError!bool {
    if (report.len == 0) return false;

    var values = std.ArrayList(i64).init(std.heap.page_allocator);
    defer values.deinit();

    var tokens = std.mem.tokenizeAny(u8, report, " ");
    while (tokens.next()) |num_str| {
        const num = try parseNumber(num_str);
        try values.append(num);
    }

    const nums = values.items;

    if (is_safe_sequence(nums)) {
        return true;
    }

    if (nums.len > 2) {
        var temp = std.ArrayList(i64).init(std.heap.page_allocator);
        defer temp.deinit();

        for (nums, 0..) |_, skip_idx| {
            temp.clearRetainingCapacity();

            for (nums, 0..) |num, i| {
                if (i != skip_idx) {
                    try temp.append(num);
                }
            }

            if (is_safe_sequence(temp.items)) {
                return true;
            }
        }
    }

    return false;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const file = try std.fs.cwd().openFile("input", .{ .mode = .read_only });
    defer file.close();

    const file_data = try file.readToEndAlloc(allocator, std.math.maxInt(usize));
    defer allocator.free(file_data);

    var lines = std.mem.tokenizeAny(u8, file_data, "\n");
    var safe_count: usize = 0;

    while (lines.next()) |line| {
        if (is_safe_with_dampener(line)) |is_valid| {
            if (is_valid) {
                safe_count += 1;
            }
        } else |err| switch (err) {
            error.InvalidFormat => continue,
            else => |e| return e,
        }
    }

    try std.io.getStdOut().writer().print("Number of safe reports with dampener: {d}\n", .{safe_count});
}
