
public class SobrecargaConstructores {

    public static void main(String[] args) {
        Persona persona1 = new Persona("Juan", 25);
        System.out.println("persona1 = " + persona1);
        
        Empleado empleado1 = new Empleado("Karla", 30, 4000);
        System.out.println("empleado1 = " + empleado1);
    }
}
