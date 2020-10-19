# example  

## file_copy_001.go  

file read  
> 1 os.Open()   
> 2 file  Close()  
> 3 file  Read([]byte)  

file write  
> 1 os.Create()  
> 2 file  Close()   
> 3 file  Write([]byte) 

## file_readwrite_001.go  

file read   
> ioutil.ReadFile   
>> output []byte    
>> string([]byte) human readable    

file write  
> ioutil.WriteFile  
>> input []byte 

##  file_write_001.go  

file write  
> os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
> file  Close()
> file  WriteString(string)


