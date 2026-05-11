/* 

=> Data Type In Go 
    
//------------------------------------------------------------------------------------------------------


=> Primitive Data Type 
	   
	1. Integer Data type 
	        
	       -> used for whole number
         
		    -> signed integer = can store positive and negative number
			   
			    int8   = 8bit 
				int16  = 16bit 
				int32  = 32bit
				int64  = 64bit 
		 
				func main(){
				
				var a int = 10 
				var b int = -5
				
				fmt.Println(a)
				fmt.Println(b)}
			
			-> unsigned integer (uint) = store only positive integer 
			     
			     var x uint = 100
	
	2. Float Types
	   
	    -> Used for decimal type 
		    
		    var pi float = 3.14

	

              func main() {
	              var price float64 = 99.99   // float64 higher precision

	              fmt.Println(price)
              }
	
	3. Boolean Type 
	      True and false 
		  
		  func main(){
		  
		     var isGoeasy = True
			 fmt.Println(isGoeasy)
			  }

// --------------------------------------------------------------------------------------------------------

=> Composite Data Type 
     
    -> Combine multiple values into one structure 

	-> can be , list of numbers , collection of users , key-value data  , custom data 

	-> Array , slice , Map , struct

// --------------------------------------------------------------------------------------------------------

=> Reference Data Type 
   
   -> These are the data type that internally refer to underlying array memory instead of storing everything directly

   -> Slice , Map , Channel , Function , Pointer  , Interface 

   1. Pointer 
      
       -> Pointer store memory address , not actual value
   
   2. Slice internally refer to underlying array 
   
   3. Map internally refer to hash table 

   4. 



*/