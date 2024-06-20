pragma circom 2.0.0;

// template input => a_b_product

/*This circuit template checks that c is the multiplication of a and b.*/  

template Example () {  

   // Declaration of signals.  
   signal input a;  
   signal input b;  
   signal output c;  
  
   c <== a * b;
   c === a_b_product;  
}

component main = Example();