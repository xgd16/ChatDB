<template>
  <div class="orbs" ref="orbsRef">
    <div v-for="n in orbCount" :key="n" class="orb" />
  </div>
  
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, ref } from 'vue'
import { gsap, Power1 } from 'gsap'

const props = withDefaults(defineProps<{
  orbCount?: number
}>(), {
  orbCount: 8
})

const orbCount = props.orbCount
const orbsRef = ref<HTMLDivElement>()
let tweens: gsap.core.Tween[] = []
let handleResize: () => void

function random(min: number, max: number) {
  return Math.random() * (max - min) + min
}

function animateOrb(el: HTMLElement) {
  const parent = el.parentElement as HTMLElement
  const w = parent.clientWidth
  const h = parent.clientHeight

  const size = random(60, 140)
  el.style.width = `${size}px`
  el.style.height = `${size}px`
  el.style.filter = `blur(${Math.floor(size / 10)}px)`
  el.style.opacity = String(random(0.35, 0.75))

  const duration = random(6000, 12000)
  const delay = random(0, 2000)

  const tween = gsap.to(el, {
    x: () => random(-w * 0.4, w * 0.4),
    y: () => random(-h * 0.4, h * 0.4),
    duration: duration / 1000,
    delay: delay / 1000,
    ease: Power1.easeInOut,
    yoyo: true,
    repeat: -1
  })
  tweens.push(tween)
}

onMounted(() => {
  const container = orbsRef.value
  if (!container) return

  const children = Array.from(container.children) as HTMLElement[]
  children.forEach((el) => {
    // 初始位置分散到容器内的随机百分比（避免一开始聚集在中间）
    const topPct = random(5, 95)
    const leftPct = random(5, 95)
    el.style.top = `${topPct}%`
    el.style.left = `${leftPct}%`
    el.style.transform = 'translate(0, 0)'

    const hue = Math.floor(random(160, 240))
    el.style.background = `radial-gradient(circle at 30% 30%, hsla(${hue}, 95%, 65%, 0.9), hsla(${hue}, 95%, 45%, 0.55) 60%, transparent 70%)`

    animateOrb(el)
  })

  handleResize = () => {
    tweens.forEach((t) => t.kill())
    tweens = []
    children.forEach((el) => animateOrb(el))
  }
  window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
  tweens.forEach((t) => t.kill())
  tweens = []
  if (handleResize) window.removeEventListener('resize', handleResize)
})
</script>

<style scoped lang="scss">
.orbs {
  position: absolute;
  inset: 0;
  overflow: hidden;
  pointer-events: none;
  z-index: 0; // 位于内容之下
}

.orb {
  position: absolute;
  border-radius: 999px;
  mix-blend-mode: screen;
  will-change: transform, opacity, filter;
}

[data-theme="dark"] .orb {
  mix-blend-mode: lighten;
}
</style>


