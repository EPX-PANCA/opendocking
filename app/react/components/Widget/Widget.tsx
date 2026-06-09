import { cn } from '@/lib/utils';
import {
  createContext,
  PropsWithChildren,
  Ref,
  useContext,
  useMemo,
} from 'react';
import { motion } from 'framer-motion';

import { useId } from '@/react/hooks/useId';

interface WidgetContextValue {
  titleId: string | undefined;
}

const Context = createContext<null | WidgetContextValue>(null);
Context.displayName = 'WidgetContext';

export function useWidgetContext() {
  const context = useContext(Context);

  if (context == null) {
    throw new Error('Should be inside a Widget component');
  }

  return context;
}

export function Widget({
  children,
  className,
  mRef,
  id,
  'aria-label': ariaLabel,
  'data-cy': dataCy,
}: PropsWithChildren<{
  className?: string;
  mRef?: Ref<HTMLDivElement>;
  id?: string;
  'aria-label'?: string;
  'data-cy'?: string;
}>) {
  const generatedId = useId();
  const titleId = ariaLabel ? undefined : `widget-title-${generatedId}`;
  const contextValue = useMemo(() => ({ titleId }), [titleId]);

  return (
    <Context.Provider value={contextValue}>
      <motion.section
        id={id}
        className={cn('widget', className)}
        ref={mRef}
        aria-label={ariaLabel}
        aria-labelledby={titleId}
        data-cy={dataCy}
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.4, ease: 'easeOut' }}
      >
        {children}
      </motion.section>
    </Context.Provider>
  );
}
