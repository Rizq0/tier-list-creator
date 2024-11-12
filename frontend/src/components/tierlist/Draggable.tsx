import React from 'react'
import { useDrag } from 'react-dnd'


const Draggable = () => {
  const [{ opacity }, dragRef] = useDrag(
    () => ({
      type: 'BOX',
      collect: (monitor) => ({
        opacity: monitor.isDragging() ? 0.4 : 1,
      }),
    }),
    [],
  )

  return (
    <div ref={dragRef} style={{ opacity }}>
      Draggable
    </div>
  )
}

export default Draggable