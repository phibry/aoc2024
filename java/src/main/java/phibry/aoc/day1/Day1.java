package phibry.aoc.day1;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class Day1 {
    public void day1() {
        try {
            var scanner = new Scanner(new File("src/main/resources/test1.txt"));
            scanner.useDelimiter("");
            while (scanner.hasNext()) {
                System.out.println(scanner.next());
                scanner.next();
            }
            scanner.close();

        } catch (FileNotFoundException e) {
        }
    }
}
