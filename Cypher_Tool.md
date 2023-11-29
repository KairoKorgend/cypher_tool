# Cypher Tool

The Cypher Tool is a command-line program that allows users to perform text encryption and decryption using different ciphers. It offers various encryption techniques and guides users through the process of securing and revealing messages. 

## Usage

To use the Cypher Tool, follow these steps:

1. Run the program by executing `./cyphertool` in your command-line environment.

2. You will be greeted with a welcome message and prompted to select an operation:
   - Enter `1` for encryption.
   - Enter `2` for decryption.

3. After selecting the operation, you will be asked to choose an encryption/decryption type:
   - Enter `1` for ROT13 encryption/decryption.
   - Enter `2` for Reverse encryption/decryption.
   - Enter `3` for Random encryption/decryption.

4. Enter the message you want to encrypt or decrypt.

5. The program will then display the result of the operation, showing the encrypted or decrypted message.

## Ciphers Used

The Cypher Tool supports the following ciphers:

1. **ROT13** (Caesar Cipher with a shift value of 13):
   - A simple letter substitution cipher that shifts each letter in the message with the letter 13 positions down the alphabet, similar to ShiftBy task.
   - Example:
     - Input: "Hello, World!"
     - ROT13-encrypted: "Uryyb, Jbeyq!"

2. **Reverse**:
   - A cipher that returns its reverse message in the Latin alphabet, similar to Reverse_Alphabet_Value task. It does not change non-alphabet characters.
   - Example:
     - Input: "Hello, World!"
     - Reverse-encrypted: "Svool, Dliow!"

3. **Random**:
   - A custom substitution cipher that maps letters from the standard alphabet to a specific set of characters.
   - Example:
     - Input: "Hello, World!"
     - Random-encrypted: "UjIIH, WHcId!"
     - Note: The character mapping is not based on a simple rule like ROT13 but is predefined by the tool.

This tool provides a convenient way to experiment with different encryption techniques and secure your messages. It can be used for educational purposes or for simple encryption and decryption needs.
