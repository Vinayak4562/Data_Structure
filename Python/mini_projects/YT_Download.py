import pytube

link = input("Enter your yt link: ")
yt = pytube.YouTube(link)
yt.streams.first().download()
print("Download",link)