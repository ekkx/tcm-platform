import { Button, DateRangePicker, Select, SelectItem } from "@heroui/react";
import type { DateValue } from "@heroui/react";
import { today } from "@internationalized/date";
import type { RangeValue } from "@react-types/shared";
import { useCallback, useRef, useState } from "react";
import { I18nProvider } from "react-aria-components";
import { CampusType, PianoType } from "~/api/pb/room/v1/room_pb";

export interface FilterValues {
  dateRange: RangeValue<DateValue> | null;
  campusType: CampusType | null;
  pianoType: PianoType | null;
}

export function Header({
  onFilter,
}: {
  onFilter?: (filters: FilterValues) => void;
}) {
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);
  const [dateRange, setDateRange] = useState<RangeValue<DateValue> | null>(
    null
  );
  const [campusType, setCampusType] = useState<CampusType | null>(null);
  const [pianoType, setPianoType] = useState<PianoType | null>(null);

  // Swipe-to-close
  const drawerRef = useRef<HTMLDivElement>(null);
  const touchStartY = useRef(0);
  const touchDeltaY = useRef(0);
  const [dragOffset, setDragOffset] = useState(0);
  const isDragging = useRef(false);

  const handleTouchStart = useCallback((e: React.TouchEvent) => {
    touchStartY.current = e.touches[0].clientY;
    touchDeltaY.current = 0;
    isDragging.current = false;
  }, []);

  const handleTouchMove = useCallback((e: React.TouchEvent) => {
    const delta = e.touches[0].clientY - touchStartY.current;
    // Only allow upward swipe (negative delta = swipe up to close)
    if (delta < 0) {
      touchDeltaY.current = delta;
      isDragging.current = true;
      setDragOffset(delta);
    } else {
      touchDeltaY.current = 0;
      setDragOffset(0);
    }
  }, []);

  const handleTouchEnd = useCallback(() => {
    if (isDragging.current && touchDeltaY.current < -80) {
      setIsDrawerOpen(false);
    }
    setDragOffset(0);
    isDragging.current = false;
  }, []);

  const handleFilter = () => {
    onFilter?.({ dateRange, campusType, pianoType });
    setIsDrawerOpen(false);
  };

  return (
    <>
      {/* Backdrop */}
      <div
        className={`fixed inset-0 z-40 bg-black/30 transition-opacity duration-300 ${
          isDrawerOpen ? "opacity-100" : "opacity-0 pointer-events-none"
        }`}
        onClick={() => setIsDrawerOpen(false)}
      />

      {/* Drawer - full width from top */}
      <div
        ref={drawerRef}
        className={`fixed top-0 inset-x-0 z-40 transition-transform duration-300 ease-in-out ${
          isDrawerOpen ? "translate-y-0" : "-translate-y-full"
        }`}
        style={
          isDragging.current && dragOffset < 0
            ? { transform: `translateY(${dragOffset}px)`, transition: "none" }
            : undefined
        }
        onTouchStart={handleTouchStart}
        onTouchMove={handleTouchMove}
        onTouchEnd={handleTouchEnd}
      >
        <div className="bg-background pt-24 pb-6 px-6 rounded-b-3xl shadow-lg">
          <div className="grid gap-4 max-w-lg mx-auto">
            <I18nProvider locale="ja">
              <DateRangePicker
                label="日付"
                labelPlacement="outside"
                fullWidth
                minValue={today("Asia/Tokyo")}
                value={dateRange}
                onChange={setDateRange}
              />
            </I18nProvider>
            <Select
              label="キャンパス"
              labelPlacement="outside"
              placeholder="すべて"
              selectedKeys={campusType !== null ? [String(campusType)] : []}
              onChange={(e) => {
                const val = e.target.value;
                setCampusType(val ? (Number(val) as CampusType) : null);
              }}
            >
              <SelectItem key={String(CampusType.NAKAMEGURO)}>
                中目黒・代官山キャンパス
              </SelectItem>
              <SelectItem key={String(CampusType.IKEBUKURO)}>
                池袋キャンパス
              </SelectItem>
            </Select>
            <Select
              label="ピアノタイプ"
              labelPlacement="outside"
              placeholder="すべて"
              selectedKeys={pianoType !== null ? [String(pianoType)] : []}
              onChange={(e) => {
                const val = e.target.value;
                setPianoType(val ? (Number(val) as PianoType) : null);
              }}
            >
              <SelectItem key={String(PianoType.GRAND)}>
                グランドピアノ
              </SelectItem>
              <SelectItem key={String(PianoType.UPRIGHT)}>
                アップライトピアノ
              </SelectItem>
              <SelectItem key={String(PianoType.NONE)}>ピアノ無</SelectItem>
            </Select>
            <Button
              fullWidth
              color="primary"
              radius="lg"
              className="mt-1"
              onPress={handleFilter}
            >
              絞り込む
            </Button>
            {/* Swipe handle */}
            <div className="flex justify-center pt-1">
              <div className="w-10 h-1 rounded-full bg-default-300" />
            </div>
          </div>
        </div>
      </div>

      {/* Header bar - always on top */}
      <div className="fixed top-6 inset-x-0 z-50 flex justify-center px-6">
        <header className="flex items-center justify-between w-full max-w-lg px-5 h-16 rounded-full bg-foreground/10 backdrop-blur-xl border-[0.5px] border-default-300">
          <div className="flex items-center gap-2">
            <span className="text-lg">🎹</span>
            <span className="text-base font-bold">TCMRSV</span>
          </div>
          <Button
            size="md"
            radius="full"
            variant="flat"
            className="text-xs text-default-500"
            onPress={() => setIsDrawerOpen(!isDrawerOpen)}
            startContent={
              <svg
                className="w-4 h-4"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <g fill="none" stroke="currentColor" strokeWidth="1.5">
                  <circle cx="11.5" cy="11.5" r="9.5" />
                  <path strokeLinecap="round" d="M18.5 18.5L22 22" />
                </g>
              </svg>
            }
          >
            絞り込み
          </Button>
        </header>
      </div>
    </>
  );
}
