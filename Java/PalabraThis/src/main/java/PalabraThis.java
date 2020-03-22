
public class PalabraThis {

    public static void main(String[] args) {
        Persona persona = new Persona("Juan");
    }
}

class Persona {

    String nombre;

    Persona(String nombre) {
        this.nombre = nombre;//this apunta a un objeto de tipo Persona
        System.out.println("impresion del operador this dentro de la clase Persona:" + this);//this apunta a un objeto de tipo persona en este momento
        Imprimir imprimir = new Imprimir();
        imprimir.imprimir(this);//this contiene una referencia al objeto Persona

    }
}

class Imprimir {

    public void imprimir(Persona p) {
        System.out.println("impresion argumento persona:" + p);//valor del objeto persona
        System.out.println("imperesion objeto actual (this)" + this);//this apunta a un objeto de tipo imprimir en este momento
    }
}
