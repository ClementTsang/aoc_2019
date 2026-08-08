import java.io.File

fun partOne(file: File) {
    var sum = 0

    file.forEachLine { line ->
        if (!line.isEmpty()) {
            val value = line.toInt()
            val fuel = value / 3 - 2

            sum += fuel
        }
    }

    println("Part 1: ${sum}")
    
}

fun partTwo(file: File) {
    var sum = 0

    file.forEachLine { line ->
        if (!line.isEmpty()) {
            var fuel = line.toInt()

            while (fuel > 0) {
                fuel = fuel / 3 - 2

                if (fuel > 0) {
                    sum += fuel
                }
            }
        }
    }

    println("Part 2: ${sum}")
}

var fileName = "input.txt"
if (args.size == 1 && !args[0].isEmpty()) {
    fileName = args[0]
}
val file = File(fileName)

partOne(file)
partTwo(file)
