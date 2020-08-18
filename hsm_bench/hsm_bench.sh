#!/bin/bash
BIN=./go-wrk-tcp.linux
ADDR=192.168.2.1:1818

set -e 

echo  金融业务指令性能测试

echo pin 加密SM1
${BIN}  -print  -t 100   -c 100  -n 100000   -p "PIN_ENC_SM1.txt"  ${ADDR}
echo pin 加密SM4
${BIN}  -print  -t 100   -c 100  -n 100000   -p "PIN_ENC_SM4.txt"  ${ADDR}

echo pin 转加密SM1
${BIN}  -print  -t 100   -c 100  -n 100000   -p "PIN_TRANS_SM1.txt"  ${ADDR}
echo pin 转加密SM4
${BIN}  -print  -t 100   -c 100  -n 100000   -p "PIN_TRANS_SM4.txt"  ${ADDR}


echo MAC生成 SM1
${BIN}  -print  -t 100   -c 100  -n 100000   -p "MAC_GEN_SM1.txt"  ${ADDR}
echo MAC生成 SM4
${BIN}  -print  -t 100   -c 100  -n 100000   -p "MAC_GEN_SM4.txt"  ${ADDR}

echo ARQC验证 SM1
${BIN}  -print  -t 100   -c 100  -n 100000   -p "ARQC_GEN_SM1.txt"  ${ADDR}
echo ARQC验证 SM4
${BIN}  -print  -t 100   -c 100  -n 100000   -p "ARQC_GEN_SM4.txt"  ${ADDR}



echo  算法性能测试

echo 随机数
${BIN}  -hex -size 4096  -t 100   -c 100  -n 1000   -p "RAND_GEN.txt"  ${ADDR}

echo SM1_ECB_ENC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM1_ECB_ENC256.txt"  ${ADDR}
echo SM1_ECB_DEC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM1_ECB_DEC256.txt"  ${ADDR}
echo SM1_ECB_ENC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM1_ECB_ENC4096.txt"  ${ADDR}
echo SM1_ECB_DEC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM1_ECB_DEC4096.txt"  ${ADDR}

echo SM1_CBC_ENC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM1_CBC_ENC256.txt"  ${ADDR}
echo SM1_CBC_DEC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM1_CBC_DEC256.txt"  ${ADDR}
echo SM1_CBC_ENC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM1_CBC_ENC4096.txt"  ${ADDR}
echo SM1_CBC_DEC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM1_CBC_DEC4096.txt"  ${ADDR}


echo SM4_ECB_ENC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM4_ECB_ENC256.txt"  ${ADDR}
echo SM4_ECB_DEC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM4_ECB_DEC256.txt"  ${ADDR}
echo SM4_ECB_ENC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM4_ECB_ENC4096.txt"  ${ADDR}
echo SM4_ECB_DEC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM4_ECB_DEC4096.txt"  ${ADDR}

echo SM4_CBC_ENC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM4_CBC_ENC256.txt"  ${ADDR}
echo SM4_CBC_DEC256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM4_CBC_DEC256.txt"  ${ADDR}
echo SM4_CBC_ENC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM4_CBC_ENC4096.txt"  ${ADDR}
echo SM4_CBC_DEC4096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM4_CBC_DEC4096.txt"  ${ADDR}


echo SM3256
${BIN}  -print -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM3256.txt"  ${ADDR}
echo SM34096
${BIN}  -hex  -size 4096  -t 100   -c 100  -n 100000   -p "SM34096.txt"  ${ADDR}


echo SM2_GEN
${BIN}  -print -hex  -size 96  -t 100   -c 100  -n 100000   -p "SM2_GEN.txt"  ${ADDR}
echo SM2_SIGN
${BIN}  -print -hex  -size 32  -t 100   -c 100  -n 100000   -p "SM2_SIGN.txt"  ${ADDR}
echo SM2_VERIFY
${BIN}  -print -hex  -size 32  -t 100   -c 100  -n 100000   -p "SM2_VERIFY.txt"  ${ADDR}

echo SM2_ENC256
${BIN}  -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM2_ENC256.txt"  ${ADDR}
echo SM2_DEC256
${BIN}  -hex  -size 256  -t 100   -c 100  -n 100000   -p "SM2_DEC256.txt"  ${ADDR}

