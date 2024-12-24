package phibry.aoc.day1;

import phibry.aoc.util.Util;

import java.util.Arrays;

public class Day1 {
    public static void day1() {
        var input = Util.readFile("src/main/resources/input1.txt");
        // var input = Util.readFile("src/main/resources/test1.txt");
        day11(input);
        day12(input);
    }

    private static void day11(int[][] input) {
        for (int[] row : input) {
            Arrays.sort(row);
        }

        var diff = 0;
        for (int i = 0; i < input[0].length; i++) {
            diff += Math.abs(input[0][i] - input[1][i]);
        }
        System.out.println(diff);
    }

    private static void day12(int[][] input) {
        // var similarityScore = 0;
        // for (int i = 0; i < input[0].length; i++) {
        //     var counter = 0;
        //     for (int j = 0; j < input[1].length; j++) {
        //         if (input[1][j] == input[0][i]) {
        //             counter++;
        //         }
        //
        //     }
        //     similarityScore += input[0][i] * counter;
        // }

        var similarityScore = 0;
        for (int item : input[0]) {
            var counter = 0;
            for (int jtem: input[1]) {
                if (item == jtem) {
                    counter++;
                }
            }
            similarityScore += item * counter;
        }

        System.out.println(similarityScore);
    }

}
