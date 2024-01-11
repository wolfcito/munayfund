import { EffectCoverflow, Pagination } from 'swiper/modules'
import { Swiper, SwiperSlide } from 'swiper/react'
import Image from 'next/image'
import 'swiper/css'
import 'swiper/css/effect-coverflow'
import 'swiper/css/pagination'
import styles from './projects.module.css'
import Link from 'next/link'
import { PROJECTS } from '@/constants'

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
          <Link href={'projects/1'}>
            <h5>{PROJECTS[1]['name-project']}</h5>
            <Image src={PROJECTS[1]['image']} alt="project" width={100} height={150} />
          </Link>
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <Link href={'projects/2'}>
            <h5>{PROJECTS[2]['name-project']}</h5>
            <Image src={PROJECTS[2]['image']} alt="project" width={100} height={150} />
          </Link>
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <Link href={'projects/3'}>
            <h5>{PROJECTS[3]['name-project']}</h5>
            <Image src={PROJECTS[3]['image']} alt="project" width={100} height={150} />
          </Link>
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <Link href={'projects/4'}>
            <h5>{PROJECTS[4]['name-project']}</h5>
            <Image src={PROJECTS[4]['image']} alt="project" width={100} height={150} />
          </Link>
        </SwiperSlide>
        <SwiperSlide className={styles.swiperSlide}>
          <Link href={'projects/5'}>
            <h5>{PROJECTS[5]['name-project']}</h5>
            <Image src={PROJECTS[5]['image']} alt="project" width={100} height={150} />
          </Link>
        </SwiperSlide>
      </Swiper>
    </div>
  )
}
