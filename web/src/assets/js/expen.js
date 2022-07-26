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
    },
    // 判断是否是移动端
    IsMobile: () => {
        return navigator.userAgent.match(/iPhone|iPad|iPod/gi) ||
            navigator.userAgent.match(/Android|BlackBerry|Opera Mini|IEMobile/gi)
    }
}
export default Expen