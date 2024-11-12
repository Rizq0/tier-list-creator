import { Flex } from '@mantine/core'
import React from 'react'
import Draggable from './Draggable'
import { DndProvider } from 'react-dnd'
import { HTML5Backend } from 'react-dnd-html5-backend'


const TierList = () => {



  return (
    <DndProvider backend={HTML5Backend}>
      <Flex justify="center" align="center" direction="column">
        <section>
          <div>Top</div>
          <div></div>
        </section>
        <Draggable />
      </Flex>
    </DndProvider>
  )
}

export default TierList