/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author meco
 */
public class HolaMundo {

    public static void main(String[] args) {
        String saludar = "Hola Mundo desde Java";
        
        System.out.println(saludar);
        System.out.println(saludar);
        System.out.println(saludar);
        
        var despedirse = "Hasta Luego";
        System.out.println(despedirse);
        
        var numero = 12;
        System.out.println(numero);
        
        var usuario = "Meco";
        var saludarte = "Hola";
        System.out.println(usuario + " " + saludarte);
        
        var i = 5;
        var j =4;
        System.out.println(i + j);
        
        var saludosIniciales = "Cómo estás?";//Al usar var de debe pasar el valor
        char saludos;//Se debe especificar el tipo de variable si no se le asigna un valor inicial
        int _dato;
        float $decimal;
        
        var nombre = "Robinson";
        var apellido = "Montes";
        System.out.println("Asociado: \n" + nombre + " " + apellido );
    }
}
