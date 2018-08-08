""" 
    Crypto-201 Generate the file
"""
import random
def main():
    """
        Generate some text 
    """
    Big = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=!@#$%^&*()_+,./;'[]<>?:\"\\|"
    for i in range(2500):
        print(random.choice(Big),end='')
main()
