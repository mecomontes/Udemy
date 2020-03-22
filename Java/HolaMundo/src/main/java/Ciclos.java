
import java.util.Scanner;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
/**
 *
 * @author meco
 */
public class Ciclos {

    public static void main(String[] args) {
        var age = 0;

        while (age < 18) {
            System.out.println("Input an age: ");
            Scanner scan = new Scanner(System.in);
            age = scan.nextInt();
            if (age >= 18) {
                System.out.println("It is an adult!");
            } else {
                System.out.println("It is a baby!");
            }
        }

        inicio:
        for (var i = 0; i < 3; i++) {

            if (i % 2 == 0) {
                break;
            } else {
                System.out.println("counter i = " + i);
                continue inicio;
            }

        }
    }

}
