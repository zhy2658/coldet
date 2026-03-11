# Coldet

3D 碰撞检测库，支持 AABB（轴对齐包围盒）、球体和点的碰撞检测。

## 功能特性

- **AABB (Axis Aligned Bounding Box)** - 轴对齐包围盒
- **Sphere** - 球体
- **Point** - 点

### 碰撞检测

| 函数 | 说明 |
|------|------|
| `CheckAabbVsAabb` | AABB 与 AABB 碰撞 |
| `CheckPointInAabb` | 点是否在 AABB 内 |
| `CheckPointInSphere` | 点是否在球体内 |
| `CheckSphereVsSphere` | 球体与球体碰撞 |
| `CheckSphereVsAabb` | 球体与 AABB 碰撞 |

### 几何计算

- `Distance(to)` - 到指定点的距离
- `ClosestPoint(to)` - 到指定点最近的表面点

## 安装

```bash
go get coldet
```

## 使用示例

```go
package main

import (
    "coldet"
)

func main() {
    // 创建 AABB
    box := coldet.NewBoundingBox([3]float32{0, 0, 0}, 2, 2, 2)
    
    // 创建球体
    sphere := coldet.NewBoundingSphere([3]float32{1, 1, 1}, 1)
    
    // 检测碰撞
    if coldet.CheckSphereVsAabb(*sphere, *box) {
        println("碰撞!")
    }
}
```

## 运行测试

```bash
go test -v ./...
```

## 依赖

- [github.com/go-gl/mathgl](https://github.com/go-gl/mathgl) - 3D 数学库
