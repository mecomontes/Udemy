
import java.util.Scanner;

public class TipoDatos {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        var usuario = scanner.nextLine();
        System.out.println("usuario = " + usuario);
        
        byte byteVar = 12;
        System.out.println("Byte Value = " + byteVar);
        
        System.out.println("bits tipo byte: " + Byte.SIZE);
        System.out.println("bytes tipo byte: " + Byte.BYTES);
        System.out.println("Minimo tipo byte: " + Byte.MIN_VALUE);
        System.out.println("Maximo tipo byte: " + Byte.MAX_VALUE);
        
        short shortVar = 12;
        System.out.println("Short Value = " + shortVar);
        
        System.out.println("bits tipo short: " + Short.SIZE);
        System.out.println("bytes tipo short: " + Short.BYTES);
        System.out.println("Minimo tipo short: " + Short.MIN_VALUE);
        System.out.println("Maximo tipo shot: " + Short.MAX_VALUE);
        
        int intVar = 12;
        System.out.println("Integer Value = " + intVar);
        
        System.out.println("bits tipo int: " + Integer.SIZE);
        System.out.println("bytes tipo int: " + Integer.BYTES);
        System.out.println("Minimo tipo int: " + Integer.MIN_VALUE);
        System.out.println("Maximo tipo int: " + Integer.MAX_VALUE);
        
        long longVar = 121234567890L;//La L indica que es una variable de tipo long
        System.out.println("long Value = " + longVar);
        
        System.out.println("bits tipo long: " + Long.SIZE);
        System.out.println("bytes tipo long: " + Long.BYTES);
        System.out.println("Minimo tipo long: " + Long.MIN_VALUE);
        System.out.println("Maximo tipo long: " + Long.MAX_VALUE);
        
        var decimal = 10;
        System.out.println("numero = " + decimal);
        
        //numero binario anteponer un 0b al numero
        var binario = 0b10;
        System.out.println("numero = " + binario);
        
        //numero octal anteponer un 0 al numero
        var octal = 010;
        System.out.println("numero = " + octal);
        
        //numero hexadecimal anteponer un 0x al numero
        var hexadecimal = 0x10;
        System.out.println("numero = " + hexadecimal);
        
        //numero flotante adicionar F al final del numero
        var flotante = 10.123F;
        System.out.println("numero = " + flotante);
        
        System.out.println("bits tipo float: " + Float.SIZE);
        System.out.println("bytes tipo float: " + Float.BYTES);
        System.out.println("Minimo tipo float: " + Float.MIN_VALUE);
        System.out.println("Maximo tipo float: " + Float.MAX_VALUE);
        
        //numero double adicionar D al final del numero
        var doble = 10.123D;
        System.out.println("numero = " + doble);
        
        System.out.println("bits tipo doble: " + Double.SIZE);
        System.out.println("bytes tipo doble: " + Double.BYTES);
        System.out.println("Minimo tipo doble: " + Double.MIN_VALUE);
        System.out.println("Maximo tipo doble: " + Double.MAX_VALUE);
        
        //Usa codificacion unicode
        char caracter = 33; // '\u0021'  o '!'
        System.out.println("numero = " + caracter);
        
        System.out.println("bits tipo caracter: " + Character.SIZE);
        System.out.println("bytes tipo caracter: " + Character.BYTES);
        System.out.println("Minimo tipo caracter: " + Character.MIN_VALUE);
        System.out.println("Maximo tipo caracter: " + Character.MAX_VALUE);
        
        boolean Booleano = true;
        System.out.println("booleano = " + Boolean.TRUE);
        System.out.println("booleano = " + Boolean.FALSE);
        System.out.println("booleano = " + Booleano);
        
        var edad = 10;
        var esAdulto = edad >= 18;
        System.out.println(esAdulto);
        
        //convertir un string en tipo entero
        var talla = Integer.parseInt("20");
        
        //convertir un string en tipo double
        var pi = Double.parseDouble("3.14159");
        
        //Tomar un caracter de una cadena iniciando con 0
        var c = "Hola".charAt(2);
        
        //Convertir una vatiable en sttring
        String texto = String.valueOf(2345);
        
        //Conversion tipo casting anteponiendo entre parentesis el tipo
        short s = 123;
        byte b = (byte) s;
    }
    
}