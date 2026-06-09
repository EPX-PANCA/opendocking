import clsx from 'clsx';
import { motion } from 'framer-motion';

import styles from './ViewLoading.module.css';

interface Props {
  message?: string;
}

export function ViewLoading({ message }: Props) {
  return (
    <motion.div
      className={clsx('row', styles.root)}
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      exit={{ opacity: 0 }}
      transition={{ duration: 0.3 }}
    >
      <div className="sk-fold">
        <div className="sk-fold-cube" />
        <div className="sk-fold-cube" />
        <div className="sk-fold-cube" />
        <div className="sk-fold-cube" />
      </div>
      {message && <span className={styles.message}>{message}</span>}
    </motion.div>
  );
}
