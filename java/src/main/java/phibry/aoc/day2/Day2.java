package phibry.aoc.day2;

import phibry.aoc.util.Util;

import java.util.ArrayList;

public class Day2 {

    public static void day2() {
        var input = Util.readFile2("src/main/resources/input2.txt");
        // var input = Util.readFile2("src/main/resources/test2.txt");
        var counter = 0;
        for (ArrayList<Integer> level : input) {
            if ((isIncreasing(level) || isDecreasing(level)) && isAdjItemAtleast1or3(level)) {
                counter++;
            }
        }
        // day21 = 624
        System.out.println(counter);
    }

    private static boolean validate() {
        return true;

    }

    private static boolean isIncreasing(ArrayList<Integer> input) {
        for (int i = 1; i < input.size(); i++) {
            if (input.get(i) < input.get(i - 1)) {
                return false;
            }
        }
        return true;
    }

    private static boolean isDecreasing(ArrayList<Integer> input) {
        for (int i = 1; i < input.size(); i++) {
            if (input.get(i) > input.get(i - 1)) {
                return false;
            }
        }
        return true;
    }

    private static boolean isAdjItemAtleast1or3(ArrayList<Integer> input) {
        for (int i = 1; i < input.size(); i++) {
            var diff = Math.abs(input.get(i - 1) - input.get(i));
            if (!(diff >= 1 && diff <= 3)) {
                return false;
            }
        }
        return true;
    }

}
