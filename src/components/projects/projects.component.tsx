import { EffectCoverflow, Pagination } from 'swiper/modules'
import { Swiper, SwiperSlide } from 'swiper/react'
import Image from 'next/image'
import 'swiper/css'
import 'swiper/css/effect-coverflow'
import 'swiper/css/pagination'
import styles from './projects.module.css'

export function Projects() {
  return (
    <div className="flex h-auto w-full">
      <Swiper
        effect={'coverflow'}
        initialSlide={1}
        grabCursor={true}
        centeredSlides={true}
        slidesPerView={3}
        coverflowEffect={{
          rotate: 50,
          stretch: -50,
          depth: 50,
          modifier: 3,
          slideShadows: true,
        }}
        pagination={true}
        modules={[EffectCoverflow]}
        className={styles.swiper}
      >
        <SwiperSlide className={styles.swiperSlide}>
          <h5>Saving the world </h5>
          <Image src="/projects/project001.jpg" alt="project" width={100} height={150} />
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <h5>Minecraft Chaos</h5>
          <Image src="/projects/project002.jpg" alt="project" width={100} height={150} />
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <h5>I am not X</h5>
          <Image src="/projects/project003.jpg" alt="project" width={100} height={150} />
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <h5>Completion </h5>
          <Image src="/projects/project004.jpg" alt="project" width={100} height={150} />
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <h5>Green Moon</h5>
          <Image src="/projects/project005.jpg" alt="project" width={100} height={150} />
        </SwiperSlide>
      </Swiper>
    </div>
  )
}
