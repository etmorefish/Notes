import binascii

from Cryptodome.Cipher import PKCS1_OAEP, AES
from Cryptodome.PublicKey import RSA
from Cryptodome.Random import get_random_bytes



data = """
{
    "code": 0,
    "data": {
        "data": [
            {
                "agency": "zq",
                "cellphone": "17355377089",
                "display_id": "2107000659",
                "faceColor": "面青",
                "faceDetectRes": "已识别",
                "faceGloss": "无光泽",
                "faceLAB": "115:138:153",
                "gender": "男",
                "id": 17433,
                "lipColor": "唇淡白",
                "lipDetectRes": "已识别",
                "lipLAB": "149:138:154",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F17355377089%2F202107%2Fface%2F20210721-075609-b52161a2-40f8-4fd4-accf-7a53205b5368.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=dN5nwZKj0NQrjMTwq5KYrSjeg3Q%3D",
                "report_url": "https://m17-dev.zhiyuntcm.com?csr=ZHIYUN&category=2#/report/SRTVBtaSOFSM0DgQ",
                "time": "2021-07-21 15:56:10"
            },
            {
                "agency": "zq",
                "cellphone": "17355377089",
                "display_id": "2107000658",
                "faceColor": "面青",
                "faceDetectRes": "已识别",
                "faceGloss": "有光泽",
                "faceLAB": "114:139:152",
                "gender": "男",
                "id": 17432,
                "lipColor": "唇淡红",
                "lipDetectRes": "已识别",
                "lipLAB": "103:148:136",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F17355377089%2F202107%2Fface%2F20210721-071601-3dad7346-cd90-4c05-8ee6-011ff6f8e082.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=nQGa%2FDvGicCx14WfP2d8WNLci%2B8%3D",
                "report_url": "https://m17-dev.zhiyuntcm.com?csr=ZHIYUN&category=2#/report/SR4r2vSEJL_6TpIw",
                "time": "2021-07-21 15:16:02"
            },
            {
                "agency": "zq",
                "cellphone": "17355377089",
                "display_id": null,
                "faceColor": "面青",
                "faceDetectRes": "已识别",
                "faceGloss": "有光泽",
                "faceLAB": "107:139:152",
                "gender": "男",
                "id": 17431,
                "lipColor": "唇淡红",
                "lipDetectRes": "已识别",
                "lipLAB": "82:153:153",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F17355377089%2F202107%2Fface%2F20210721-070107-f4cb422d-e202-4263-aeab-0f89e254520e.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=5zqK3GamASHYh02qK%2BcESt5mF7o%3D",
                "report_url": null,
                "time": "2021-07-21 15:01:08"
            },
            {
                "agency": "zq",
                "cellphone": "17355377089",
                "display_id": "2107000657",
                "faceColor": "面青",
                "faceDetectRes": "已识别",
                "faceGloss": "少光泽",
                "faceLAB": "118:138:152",
                "gender": "男",
                "id": 17430,
                "lipColor": "唇淡红",
                "lipDetectRes": "已识别",
                "lipLAB": "87:152:153",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F17355377089%2F202107%2Fface%2F20210721-061705-d873b0c8-e7eb-4958-97dc-45ecabf5a7e1.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=4RHYSODDv0ZJfs6Z3xwNDkBQBrw%3D",
                "report_url": "https://m17-dev.zhiyuntcm.com?csr=ZHIYUN&category=2#/report/SRb5KKg062NIZNug",
                "time": "2021-07-21 14:17:05"
            },
            {
                "agency": "zq",
                "cellphone": "17355377089",
                "display_id": "2107003544",
                "faceColor": "面青",
                "faceDetectRes": "已识别",
                "faceGloss": "有光泽",
                "faceLAB": "110:139:154",
                "gender": "男",
                "id": 17429,
                "lipColor": "唇淡红",
                "lipDetectRes": "已识别",
                "lipLAB": "103:148:136",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F17355377089%2F202107%2Fface%2F20210721-061220-c1ad8cab-6e8b-4b3c-920a-e2b3cf09d55b.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=lTwS1QUiwzihjV%2Fcr%2B2uqMW%2B034%3D",
                "report_url": "https://m17-dev.zhiyuntcm.com/#/report/SR0p2eN9o6GkgU_w",
                "time": "2021-07-21 14:12:21"
            },
            {
                "agency": "zq",
                "cellphone": "18595438423",
                "display_id": "2107000656",
                "faceColor": "面白",
                "faceDetectRes": "已识别",
                "faceGloss": "有光泽",
                "faceLAB": "190:132:134",
                "gender": "男",
                "id": 17428,
                "lipColor": "唇淡白",
                "lipDetectRes": "已识别",
                "lipLAB": "155:144:138",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F18595438423%2F202107%2Fface%2F20210721-023123-de845372-c60e-4181-aac5-b0aa0a2376bb.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=pupSrDmpwrE7dYQDkz8IuXGz3Vk%3D",
                "report_url": "https://m17-dev.zhiyuntcm.com?csr=ZHIYUN&category=2#/report/SRvXQp7TKgaOlckw",
                "time": "2021-07-21 10:31:24"
            },
            {
                "agency": "zq",
                "cellphone": "18595438423",
                "display_id": "2107000655",
                "faceColor": "面白",
                "faceDetectRes": "已识别",
                "faceGloss": "少光泽",
                "faceLAB": "185:134:135",
                "gender": "男",
                "id": 17427,
                "lipColor": "唇淡白",
                "lipDetectRes": "已识别",
                "lipLAB": "163:138:144",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F18595438423%2F202107%2Fface%2F20210721-021455-b961ec94-8aba-4d0c-a903-e1f74792ce2e.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=SzJe6Hqiok5MxW3pt0nE9hTybHY%3D",
                "report_url": "https://m17-dev.zhiyuntcm.com?csr=ZHIYUN&category=2#/report/SRCQT_ruJTcdLSzg",
                "time": "2021-07-21 10:14:56"
            },
            {
                "agency": "zq",
                "cellphone": "18595438423",
                "display_id": "2107000654",
                "faceColor": "面白",
                "faceDetectRes": "已识别",
                "faceGloss": "有光泽",
                "faceLAB": "182:133:135",
                "gender": "男",
                "id": 17426,
                "lipColor": "唇淡白",
                "lipDetectRes": "已识别",
                "lipLAB": "155:141:142",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/zq%2F18595438423%2F202107%2Fface%2F20210721-021054-34210793-c6c7-42f0-a811-719c8e0b047d.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=jCC8TSB3ohqUeI67ykrPIPU85aE%3D",
                "report_url": "https://m17-dev.zhiyuntcm.com?csr=ZHIYUN&category=2#/report/SRnswfRf4pHwr_cg",
                "time": "2021-07-21 10:10:55"
            },
            {
                "agency": "wu123",
                "cellphone": "15690723529",
                "display_id": null,
                "faceColor": "正常",
                "faceDetectRes": "已识别",
                "faceGloss": "少光泽",
                "faceLAB": "150:137:135",
                "gender": "男",
                "id": 17425,
                "lipColor": "唇淡红",
                "lipDetectRes": "已识别",
                "lipLAB": "107:147:135",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/wu123%2F15690723529%2F202107%2Fface%2F20210721-013752-7dd148f6-a9bf-470e-af43-2bb1cdeaeb7c.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=pBN5hdMLJYq2K%2BkVqf33kqCaiRg%3D",
                "report_url": null,
                "time": "2021-07-21 09:37:53"
            },
            {
                "agency": "wu123",
                "cellphone": "15690723529",
                "display_id": null,
                "faceColor": "正常",
                "faceDetectRes": "已识别",
                "faceGloss": "少光泽",
                "faceLAB": "148:136:137",
                "gender": "男",
                "id": 17424,
                "lipColor": "唇淡红",
                "lipDetectRes": "已识别",
                "lipLAB": "102:146:136",
                "photo_url": "http://zy-test-photo.oss-cn-shanghai.aliyuncs.com/wu123%2F15690723529%2F202107%2Fface%2F20210721-013357-4b0ad51c-5168-4044-b38b-b8ebf069e5c5.jpg?OSSAccessKeyId=LTAIM1qpydZVBBzg&Expires=1626858288&Signature=iN78NmykxCaOk6m9dTG%2FdaKfjWo%3D",
                "report_url": null,
                "time": "2021-07-21 09:33:58"
            }
        ],
        "total": 17313
    },
    "message": ""
}"""


def generate_key():
    """
    生成密钥对
    :return:
    """
    key = RSA.generate(2048)
    private_key = key.export_key()
    print(private_key)
    with open('prikey.pem', 'wb') as f:
        f.write(private_key)

    public_key = key.publickey().export_key()
    print(public_key)
    with open('pubkey.pem', 'wb') as f:
        f.write(public_key)

generate_key()


def rsa_encrypt(key):
    """
    加密
    :return:
    """
    public_key = RSA.import_key(open('pubkey.pem', 'r').read())
    # data = data.encode()
    cipher = PKCS1_OAEP.new(public_key)
    key = cipher.encrypt(key)
    print("加密后的key", key)
    # print(binascii.b2a_hex(msg))
    return key

key = get_random_bytes(16)
# 16个字节的密码
key = b'1234567890123456'
rsa_encrypt(key)


def aes_encrypt(data):
    cipher = AES.new(key, AES.MODE_EAX)
    data = data.encode()
    ciphertext, tag = cipher.encrypt_and_digest(data)
    # print(cipher.nonce, tag, ciphertext)
    file_out = open("encrypted", "wb")
    [file_out.write(x) for x in (cipher.nonce, tag, ciphertext)]
    file_out.close()
    return cipher.nonce, tag, ciphertext


print(aes_encrypt(data))



