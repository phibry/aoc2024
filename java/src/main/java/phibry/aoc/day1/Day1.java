package phibry.aoc.day1;

import phibry.aoc.util.Util;

import java.util.Arrays;

public class Day1 {
    public void day1() {
        var input = Util.readFile("src/main/resources/test1.txt");
        for (int[] row : input) {
            Arrays.sort(row);
            for (int i : row) {
                System.out.print(i + " ");
            }
            System.out.println();
        }
    }
}
