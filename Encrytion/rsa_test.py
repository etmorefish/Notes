import binascii

from Cryptodome.PublicKey import RSA
from Cryptodome.Cipher import PKCS1_OAEP

def generate_key():
    """
    生成密钥对
    :return:
    """
    key = RSA.generate(1024)
    private_key = key.export_key()
    print(private_key)
    with open('prikey.pem', 'wb') as f:
        f.write(private_key)

    public_key = key.publickey().export_key()
    print(public_key)
    with open('pubkey.pem', 'wb') as f:
        f.write(public_key)

generate_key()

def encrypt():
    """
    加密
    :return:
    """
    public_key = RSA.import_key(open('pubkey.pem', 'r').read())
    # print(public_key.e)
    # print(public_key.n)
    # print(len(str(public_key.n)))
    # print(public_key.d)
    data = '加密通话'.encode()
    cipher = PKCS1_OAEP.new(public_key)
    msg = cipher.encrypt(data)
    # print(msg)
    # print(binascii.b2a_hex(msg))
    return msg

# encrypt()

def decrypt():
    """
    解密
    :return:
    """
    private_key = RSA.import_key(open('prikey.pem').read())
    msg = encrypt()
    print(f'加密信息 {msg}')
    cipher = PKCS1_OAEP.new(private_key)
    res = cipher.decrypt(msg).decode('utf-8')
    return res


print(decrypt())
