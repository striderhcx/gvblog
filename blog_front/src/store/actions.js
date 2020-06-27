// import {UpdateStudentInfo} from './mutation_defind'

// export default {
    // 异步操作
    // 1. 不优雅的写法１
    // updateInfoAsync(context, pay_load) {
    //     // context 可以理解为ｓｔｏｒｅ对象
    //     setTimeout(() => {
    //         context.commit(UpdateStudentInfo)
    //         console.log("异步操作完成！")
    //         console.log("updateInfoAsync", pay_load)
    //         pay_load.onSuccess()
    //     }, 1000)
    // }

    
    // 2. 优雅写法２
    // updateInfoAsync(context, pay_load) {
    //     return new Promise((resolve, reject) => {
    //         setTimeout(() => {
    //             context.commit(UpdateStudentInfo)
    //             console.log("异步操作完成！")
    //             console.log("updateInfoAsync", pay_load)
    //             resolve('2222222')
    //         }, 1000)
    //     })
    // }
// }