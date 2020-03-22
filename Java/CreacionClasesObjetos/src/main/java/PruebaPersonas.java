public class PruebaPersonas {

    public static void main(String[] args) {
        //Creacion de un objeto tipo persona

        //definicion de la variable tipo objeto
        Persona persona1;

        //Instanciando (creando) un objeto de la clase persona
        persona1 = new Persona();

        //Accedemos al objeto persona y llamamos al metodo desplegar Nombres
        persona1.desplegarNombres();

        // Modificamos los valores de los atributos del objeto persona
        persona1.nombre = "Meco";
        persona1.apellido = "Montes";

        //Volvemos a imprimir los valores 
        System.out.println("");
        persona1.desplegarNombres();

        //CRear un segundo objeto tipo persona 
        Persona persona2 = new Persona();
        persona2.nombre = "Katthy";
        persona2.apellido = "Perez";
        System.out.println("");
        persona2.desplegarNombres();
    }
}
