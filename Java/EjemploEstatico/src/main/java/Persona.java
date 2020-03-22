public class Persona {
    
    private int idPersona;//se asocia con un objeto (instancia)
    private String nombre;//se asocia con un objeto (instancia)
    private static int contadorPersonas;//se asocia con la clase
    
    public Persona(String nombre){
        this.idPersona = ++contadorPersonas;
        this.nombre = nombre;
    }
    
    public int getIdPersona(){
        return this.idPersona;
    }
    
//    public void setIdPersona(int idPersona){
//        this.idPersona = idPersona;
//    }
    
    public String getNombre(){
        return this.nombre;
    }
    
    public void setNombre(String nombre){
        this.nombre = nombre;
    }
    
    public static int getContadorPersonas(){
        return contadorPersonas;
    }
    
    @Override
    public String toString(){
        return "idPersona: " + idPersona + ", nombre: " + nombre + ", contadorPersonas:" + contadorPersonas; 
    }
}
