//Destruction
Demo :: Demo (int x ): x( x) {
cout << " cons :" << x << " ␣";
}

class Cont {
private :
Demo d;
public :
Cont (int x );
};
Cont :: Cont (int x ): d( x) {}
int main ()
{
Cont c (100);
return 0;
}




//Type to another Type

class Animal
2
{
3 public :
4 virtual void f () {}
5 };
6 class Dog : public Animal {};
7 class Cat : public Animal {};
8
9 int main ()
10
{
11 Animal *a = new Dog () , *b = new Cat ();
12 Dog *x = dynamic_cast < Dog *> (a),
13 *y = dynamic_cast < Dog *> (b );
14 cout << x << " ,␣" << y;
15 return 0;
16
}


//MODE
class B: мод virtual
A
{
...
};
//Здесь «мод» – это public, protected или private.

//Extends

//----1-----
class Animal
2
{
3 private :
4 string species ;
5 public :
6 Animal ( const string & species );
7 };
8
9 Animal :: Animal ( const string & species ):
10 species ( species ) {}
11
12 class Dog : public Animal
13
{
14 private :
15 string breed ;
16 public :
17 Dog ( const string & bread );
18 };
19
20 Dog :: Dog ( const string & bread ):
21 Animal (" Canis ␣ lupus ␣ familiaris "), breed ( breed ) {}


//-----2-----
class R
2
{
3 public :
4 int q;
5 R () { q = 13; }
6 };
7
8 class A: virtual R { public : void setq (int x ); };
9 class B: virtual R { public : int getq (); };
10 class C: public A , public B { };
11
12 void A :: setq ( int x) { q = x; }
13 int B :: getq () { return q; }
14
15 int main ( void
)
16
{
17 C c;
18 c . setq (666);
19 cout << c . getq (); // Выведет 666
20 return 0;
21
}
