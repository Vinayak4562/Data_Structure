# with open("file.txt","w") as file:
#     file.write("Writing files is great\n")
#     file.write("Here's another line of txt\n")
#     file.write("Let's go check it out\n")

# a - appends to, preserving orignal contents
# NO CONTROL OVER CURSOR
# with open("FILE.txt","a") as file:
#     file.seek(0)
#     file.write("APPENDIG LATER!!!")

# r+ 
# with open("file.txt", "r+")as file:
#     file.write(":)")
#     file.seek(10)
#     file.write(":(")

# r+ will not create a file if it doesn't exist
with open("hello.txt", "a") as file:
    file.write("HELLO!!!")
