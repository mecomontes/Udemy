
public class PalabraReturnTiposPrimitivos {

    public static void main(String[] args) {
        
        sumarSinRetornarValor(3, 6);
        
        int resultado = sumarRetornandoValor(4,2);
        System.out.println("resultado sumar retornando valor= " + resultado);
    }

    private static void sumarSinRetornarValor(int a, int b) {
        System.out.println("el resultado sumar sin retornar valor:" + (a + b));
        return;//opcional
    }
    
    private static int sumarRetornandoValor(int a, int b){
        if(a == 0 && b == 0){
            System.out.println("Debe proporcionar valores distintos de cero");
            return 0;
        }
        System.out.println("los valores proporcionados son distintos de cero");
        return a + b;
    }
    
    
    
}
