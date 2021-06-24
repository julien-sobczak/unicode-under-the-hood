import java.util.ArrayList;
import java.util.List;

public class BenchmarkString {
    public static void main(String[] args) {
        List<String> results = new ArrayList<>(); // Keep strings to avoid GC
        Runtime runtime = Runtime.getRuntime();

        long startTime = System.nanoTime();
        long memoryBefore = runtime.totalMemory() - runtime.freeMemory();

        String loremIpsum = """
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur.
Excepteur sint occaecat cupidatat non proident,
sunt in culpa qui officia deserunt mollit anim id est laborum.
        """;

        // StringLatin1
        for (int i = 0; i < 1000000; i++) {
            results.add((loremIpsum + i).toLowerCase());
        }

        long timeElapsed = System.nanoTime() - startTime;
        long memoryAfter = runtime.totalMemory() - runtime.freeMemory();
        long memoryUsed = memoryAfter - memoryBefore;
        System.out.println("Execution time: " + (timeElapsed / 1000000) + "ms");
        System.out.println("Memory usage: " +  (memoryUsed / 10000000) + "MB" );

        // StringUTF16
        for (int i = 0; i < 1000000; i++) {
            results.add((loremIpsum + "😀" + i).toLowerCase());
        }

        timeElapsed = System.nanoTime() - startTime;
        memoryAfter = runtime.totalMemory() - runtime.freeMemory();
        memoryUsed = memoryAfter - memoryBefore;
        System.out.println("Execution time: " + (timeElapsed / 1000000) + "ms");
        System.out.println("Memory usage: " +  (memoryUsed / 10000000) + "MB" );
    }
}
