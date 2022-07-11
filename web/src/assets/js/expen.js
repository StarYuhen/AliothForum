import copy from 'copy-to-clipboard'

const Expen = {
    // 复制内容到剪切板
    CopySrc: async (src) => {
        copy(src)
    },
    // 将登陆的值存进缓存中
    InterFaceLogin: async (data) => {
        localStorage.jwt = data.Jwt
        localStorage.name = data.Name
        localStorage.img = data.ImgUrl
    }
}
export default Expen