
public class PalabraNull {

    public static void main(String[] args) {
        Persona persona1 = new Persona("Juan");
        System.out.println("persona1:" + persona1.obtenerNombre());
        
        Persona persona2 = persona1;
        System.out.println("persona2:" + persona2.obtenerNombre());
        
        persona1 = null;
        if(persona1 != null)
            System.out.println("persona1:" + persona1.obtenerNombre());
        else
            System.out.println("la variable persona1 no apunta a ningun objeto");
        
        persona2 = null;
        
        System.gc();//garbage collector solo planifica la llamada
        
    }
}

class Persona{
    String nombre;
    
    Persona(String nombre){
        this.nombre = nombre;
    }
    
    public String obtenerNombre(){
        return this.nombre;
    }
}