package libgralloc_priv

import (
    "android/soong/android"
    "android/soong/cc"
    "fmt"
    "strings"
)

func init() {
    //该打印会在执行mm命令时，打印在屏幕上
    fmt.Println("libgralloc_priv want to conditional Compile")
    android.RegisterModuleType("cc_libgralloc_priv", DefaultsFactory)
}

func DefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, Defaults)
    return module
}

func Defaults(ctx android.LoadHookContext) {
    type props struct {
        Cflags []string
        Include_dirs []string
    }
    p := &props{}
    p.Cflags = globalCflagsDefaults(ctx)
    p.Include_dirs = globalIncludeDefaults(ctx)
    ctx.AppendProperties(p)
}

func globalIncludeDefaults(ctx android.BaseContext) ([]string) {
    var include_dirs []string
    if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali-tDVx")) {
        include_dirs = append(include_dirs,"hardware/rockchip/libgralloc/bifrost")
    } else if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali-t860") || strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali-t760")) {
        include_dirs = append(include_dirs,"hardware/rockchip/libgralloc/midgard")
    } else if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali400") || strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali450")) {
        include_dirs = append(include_dirs,"hardware/rockchip/libgralloc/utgard")
    } else if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"G6110")) {
        fmt.Println("G6110 don't contains hardware/rockchip/libgralloc!");
    } else {
        include_dirs = append(include_dirs,"hardware/rockchip/libgralloc")
    }
    return include_dirs

}

func globalCflagsDefaults(ctx android.BaseContext) ([]string) {
    var cppflags []string
    //该打印输出为: TARGET_PRODUCT:rk3328 fmt.Println("TARGET_PRODUCT:",ctx.AConfig().Getenv("TARGET_PRODUCT")) //通过 strings.EqualFold 比较字符串，可参考go语言字符串对比
    if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali-t720")) {
        //添加 DEBUG 宏定义
        cppflags = append(cppflags,"-DMALI_PRODUCT_ID_T72X=1")
        cppflags = append(cppflags,"-DMALI_AFBC_GRALLOC=0")
    } else if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali-t760")) {
        cppflags = append(cppflags,"-DMALI_PRODUCT_ID_T76X=1")
        cppflags = append(cppflags,"-DMALI_AFBC_GRALLOC=1")
    } else if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"mali-t860")) {
        cppflags = append(cppflags,"-DMALI_PRODUCT_ID_T86X=1")
        cppflags = append(cppflags,"-DMALI_AFBC_GRALLOC=1")
    } else if (strings.EqualFold(ctx.AConfig().Getenv("TARGET_BOARD_PLATFORM_GPU"),"G6110")) {
        cppflags = append(cppflags,"-DGPU_G6110")
    }
    //将需要区分的环境变量在此区域添加 //....
    return cppflags
}