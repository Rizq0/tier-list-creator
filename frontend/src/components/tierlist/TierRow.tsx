import React, { useState, FC } from 'react'
import { ItemTypes } from '../../utils/ItemTypes';

export interface TierRowProps {
  children: React.ReactNode
}

const TierRow: FC<TierRowProps> = ({ children }) => {
  const [hasDopped, setHasDropped] = useState(false);
  const [lastDroppedColor, setLastDroppedColor] = useState(false);

  const [{ isOver, isOverCurrent }, drop] = useDrop(
    () => ({
      accept: ItemTypes.BOX,
      drop(_item: unknown, monitor) {
        setHasDropped(true)
        setHasDroppedOnChild(didDrop)
      },
      collect: (monitor) => ({
        isOver: monitor.isOver(),
        isOverCurrent: monitor.isOver({ shallow: true }),
      }),
    }),
    [greedy, setHasDropped, setHasDroppedOnChild],
  )


  return (
    <div>TierRow</div>
  )
}

export default TierRow