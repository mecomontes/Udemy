
public class PasoPorReferencia {

    public static void main(String[] args) {
        Persona persona = new Persona();
        persona.cambiarNombre("Juan");

        System.out.println("valor Nombre: " + persona.obtenerNombre());

        modificarPersona(persona);
        
        System.out.println("valor nuevo nombre: " + persona.obtenerNombre());

    }

    private static void modificarPersona(Persona personaArg) {
        personaArg.cambiarNombre("Carlos");
    }
}
