import { Container } from '@mantine/core'
import React from 'react'

const MaxWidth = ({ children } : { children: React.ReactNode }) => {
  return (
    <Container
      maw={{ xs: '100%', xl: 1340 }}
      px={{ xs: 'md', sm: 'lg' }}
    >
      {children}
    </Container>
  )
}

export default MaxWidth