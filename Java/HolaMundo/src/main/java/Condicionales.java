
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
public class Condicionales {

    public static void main(String[] args) {

        //operadores ternarios: una simplificacion del if-else
        // variable = expresion ? if: else;
        var resultado = (1 < 2) ? "Verdaadero" : "Falso";
        System.out.println("resultado: " + resultado);

        System.out.println("Input age: ");
        Scanner scanner = new Scanner(System.in);
        int edad = scanner.nextInt();

        if (edad > 18) {
            System.out.println("Eres un adulto");

        } else if (edad == 18) {
            System.out.println("Es hora de pedir tu cedula");
        } else {
            System.out.println("eres menor de edad");
        }

        String season;
        System.out.println("Input a Month: ");
        Scanner scanner1 = new Scanner(System.in);
        String month = scanner1.nextLine();

        if ("December".equals(month) || "January".equals(month) || "Febrary".equals(month)) {
            season = "Winter";
        } else if ("March".equals(month) || "April".equals(month) || "May".equals(month)) {
            season = "Spring";
        } else if ("june".equals(month) || "july".equals(month) || "Agust".equals(month)) {
            season = "Summer";
        } else {
            season = "Fall";
        }
        System.out.println("We are on " + season);

        Scanner scanner2 = new Scanner(System.in);
        int day = scanner2.nextInt();

        switch (day) {
            case 1:
                System.out.println("Lunes");
                break;
            case 2:
                System.out.println("Martes");
                break;
            case 3:
                System.out.println("Miercoles");
                break;
            case 4:
                System.out.println("Jueves");
                break;
            case 5:
                System.out.println("Viernes");
                break;
            case 6:
                System.out.println("Sabado");
                break;
            case 7:
                System.out.println("Domingo");
                break;
            default:
                System.out.println("It is not a day");

        }

    }

}
