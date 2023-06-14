public class Main {
    public static void main(String[] args) {
        String s;

        s="template < " +
                "template <  < typename T , int S > " +
                "class C , " +
                "typename T , int N " +
                ">";
       // System.out.println(s);
        System.out.println(new Parser(s).parse());

        System.out.println("ok");
        s="template < class C , typename T , int N >";
        //System.out.println(s);
        System.out.println(new Parser(s).parse());

    }
}