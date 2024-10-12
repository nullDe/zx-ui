# 创建目录
mkdir -p components core docs hooks theme utils

# 进入每个目录并执行 pnpm init
for i in components core docs hooks  theme utils; do
 cd $i
 pnpm init
 cd ..
done