package phibry.aoc.util;

import java.io.File;
import java.util.*;
import java.util.stream.IntStream;

public class Util {
    public static int[][] readFile(String filename) {
        try {
            var scanner = new Scanner(new File(filename));
            var matrix = new ArrayList<ArrayList<Integer>>();
            while (scanner.hasNextLine()) {
                var line = scanner.nextLine();
                String[] parts = line.split("\\s+");
                matrix.add(new ArrayList<>(Arrays.asList(Integer.parseInt(parts[0]), Integer.parseInt(parts[1]))));
            }
            scanner.close();
            return IntStream.range(0, matrix.get(0).size())
                    .mapToObj(row -> matrix.stream()
                            .mapToInt(inner -> inner.get(row))
                            .toArray())
                    .toArray(int[][]::new);
        } catch (Exception e) {
            System.out.println(e);
            // TODO: handle exception
        }

        return null;
    }
}
